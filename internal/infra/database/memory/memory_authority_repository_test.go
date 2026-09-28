package memory

import (
	"context"
	"errors"
	"go-auth-service/internal/domain"
	"go-auth-service/internal/domain/authority"
	"testing"
)

func newTestAuthorityMemoryRepository(t *testing.T) (*MemoryAuthorityRepository, context.Context) {
	t.Helper()

	repo, err := New(make(map[int64]*authority.Authority))
	if err != nil {
		t.Fatalf("failed to create authority repository: %v", err)
	}

	return repo, context.Background()
}

func TestMemoryAuthorityRepository_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectedName := "authority.create"
		expectedID := int64(1)

		repo, ctx := newTestAuthorityMemoryRepository(t)
		result, err := repo.Create(
			ctx, authority.Authority{Name: expectedName},
		)

		if err != nil {
			t.Fatalf("failed to create authority: %v", err)
		}
		if result.ID != expectedID {
			t.Error("created data ID not as expected")
		}
		if result.Name != expectedName {
			t.Error("created data not as expected")
		}
	})

	t.Run("sequential", func(t *testing.T) {
		sequentialID := int64(1)
		repo, ctx := newTestAuthorityMemoryRepository(t)
		sequentialAuthority, err := repo.Create(
			ctx, authority.Authority{Name: "new.authority"},
		)

		if err != nil {
			t.Fatalf("failed to create second authority: %v", err)
		}

		if sequentialAuthority.ID != sequentialID {
			t.Errorf(
				"expected sequential authority ID %d, got %d",
				sequentialID, sequentialAuthority.ID,
			)
		}
		sequentialID++

		if err := repo.DeleteByID(ctx, sequentialAuthority.ID); err != nil {
			t.Fatalf("unexpected error while delete authority: %v", err)
		}

		newAuthority, err := repo.Create(
			ctx, authority.Authority{Name: "new.authority"},
		)

		if err != nil {
			t.Fatalf("failed to create authority after delete: %v", err)
		}

		if newAuthority.ID != sequentialID {
			t.Errorf("failed to create sequential ID after delete: %v", err)
		}

	})

	t.Run("duplicate", func(t *testing.T) {
		authorityName := "authority.create"
		repo, ctx := newTestAuthorityMemoryRepository(t)
		repo.Create(
			ctx, authority.Authority{Name: authorityName},
		)

		_, err := repo.Create(
			ctx, authority.Authority{Name: authorityName},
		)

		var appErr *domain.AppError
		if !errors.As(err, &appErr) {
			t.Fatalf("expected AppError, got %T", err)
		}

		if appErr.Cause != domain.ErrDuplicate {
			t.Errorf("expected error duplicate, got %T", err)
		}
	})

	t.Run("isolated", func(t *testing.T) {
		createdName := "authority.create"
		repo, ctx := newTestAuthorityMemoryRepository(t)
		created, _ := repo.Create(
			ctx, authority.Authority{Name: createdName},
		)

		created.Name = "updated"
		found, err := repo.FindByID(ctx, created.ID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if found.Name != createdName {
			t.Errorf(
				"repository entity was modified through returned entity: got %q",
				found.Name,
			)
		}
	})
}

func TestMemoryAuthorityRepository_FindByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo, ctx := newTestAuthorityMemoryRepository(t)
		created, err := repo.Create(ctx, authority.Authority{Name: "authority.find"})
		if err != nil {
			t.Fatalf("failed to create authority: %v", err)
		}

		found, err := repo.FindByID(ctx, created.ID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if found.ID != created.ID {
			t.Errorf("expected ID %d, got %d", created.ID, found.ID)
		}
		if found.Name != created.Name {
			t.Errorf("expected name %q, got %q", created.Name, found.Name)
		}
	})

	t.Run("not found", func(t *testing.T) {
		repo, ctx := newTestAuthorityMemoryRepository(t)

		_, err := repo.FindByID(ctx, 999)
		if err == nil {
			t.Fatal("expected error for non-existent ID, got nil")
		}

		var appErr *domain.AppError
		if !errors.As(err, &appErr) {
			t.Fatalf("expected AppError, got %T", err)
		}

		if appErr.Cause != domain.ErrNotFound {
			t.Errorf("expected ErrNotFound, got %v", appErr.Cause)
		}
	})

	t.Run("returned entity is copy", func(t *testing.T) {
		repo, ctx := newTestAuthorityMemoryRepository(t)
		created, _ := repo.Create(ctx, authority.Authority{Name: "authority.find"})

		found, err := repo.FindByID(ctx, created.ID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		found.Name = "modified.name"

		again, err := repo.FindByID(ctx, created.ID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if again.Name != "authority.find" {
			t.Errorf(
				"modifying returned entity affected repository: got %q",
				again.Name,
			)
		}
	})
}

func TestMemoryAuthorityRepository_FindByIDs(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo, ctx := newTestAuthorityMemoryRepository(t)
		a1, _ := repo.Create(ctx, authority.Authority{Name: "authority.a"})
		a2, _ := repo.Create(ctx, authority.Authority{Name: "authority.b"})

		results, err := repo.FindByIDs(ctx, []int64{a1.ID, a2.ID})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(results) != 2 {
			t.Fatalf("expected 2 results, got %d", len(results))
		}

		if results[0].ID != a1.ID || results[0].Name != a1.Name {
			t.Errorf("first result mismatch")
		}
		if results[1].ID != a2.ID || results[1].Name != a2.Name {
			t.Errorf("second result mismatch")
		}
	})

	t.Run("one ID missing", func(t *testing.T) {
		repo, ctx := newTestAuthorityMemoryRepository(t)
		a1, _ := repo.Create(ctx, authority.Authority{Name: "authority.a"})

		_, err := repo.FindByIDs(ctx, []int64{a1.ID, 999})
		if err == nil {
			t.Fatal("expected error for missing ID, got nil")
		}

		var appErr *domain.AppError
		if !errors.As(err, &appErr) {
			t.Fatalf("expected AppError, got %T", err)
		}

		if appErr.Cause != domain.ErrNotFound {
			t.Errorf("expected ErrNotFound, got %v", appErr.Cause)
		}
	})

	t.Run("returned entities are copies", func(t *testing.T) {
		repo, ctx := newTestAuthorityMemoryRepository(t)
		a1, _ := repo.Create(ctx, authority.Authority{Name: "authority.a"})
		a2, _ := repo.Create(ctx, authority.Authority{Name: "authority.b"})

		results, err := repo.FindByIDs(ctx, []int64{a1.ID, a2.ID})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		results[0].Name = "modified.a"
		results[1].Name = "modified.b"

		again, err := repo.FindByIDs(ctx, []int64{a1.ID, a2.ID})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if again[0].Name != "authority.a" {
			t.Errorf("modifying result[0] affected repo: got %q", again[0].Name)
		}
		if again[1].Name != "authority.b" {
			t.Errorf("modifying result[1] affected repo: got %q", again[1].Name)
		}
	})
}

func TestMemoryAuthorityRepository_ExistByName(t *testing.T) {
	t.Run("exists", func(t *testing.T) {
		repo, ctx := newTestAuthorityMemoryRepository(t)
		repo.Create(ctx, authority.Authority{Name: "authority.exists"})

		if !repo.ExistByName("authority.exists") {
			t.Error("expected ExistByName to return true for existing name")
		}
	})

	t.Run("doesn't exist", func(t *testing.T) {
		repo, _ := newTestAuthorityMemoryRepository(t)

		if repo.ExistByName("nonexistent") {
			t.Error("expected ExistByName to return false for non-existent name")
		}
	})
}

func TestMemoryAuthorityRepository_Update(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo, ctx := newTestAuthorityMemoryRepository(t)
		created, _ := repo.Create(ctx, authority.Authority{Name: "authority.update"})

		input := authority.Authority{
			ID:   created.ID,
			Name: "authority.updated",
		}

		updated, err := repo.Update(ctx, input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if updated.ID != created.ID {
			t.Errorf("expected ID %d, got %d", created.ID, updated.ID)
		}
		if updated.Name != "authority.updated" {
			t.Errorf("expected name %q, got %q", "authority.updated", updated.Name)
		}
	})

	t.Run("not found", func(t *testing.T) {
		repo, ctx := newTestAuthorityMemoryRepository(t)

		input := authority.Authority{
			ID:   999,
			Name: "authority.update",
		}

		_, err := repo.Update(ctx, input)
		if err == nil {
			t.Fatal("expected error for non-existent ID, got nil")
		}

		var appErr *domain.AppError
		if !errors.As(err, &appErr) {
			t.Fatalf("expected AppError, got %T", err)
		}

		if appErr.Cause != domain.ErrNotFound {
			t.Errorf("expected ErrNotFound, got %v", appErr.Cause)
		}
	})

	t.Run("same name", func(t *testing.T) {
		repo, ctx := newTestAuthorityMemoryRepository(t)
		created, _ := repo.Create(ctx, authority.Authority{Name: "authority.update"})

		input := authority.Authority{
			ID:   created.ID,
			Name: created.Name,
		}

		updated, err := repo.Update(ctx, input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if updated.Name != created.Name {
			t.Errorf("expected name %q, got %q", created.Name, updated.Name)
		}
	})

	t.Run("rename", func(t *testing.T) {
		repo, ctx := newTestAuthorityMemoryRepository(t)
		created, _ := repo.Create(ctx, authority.Authority{Name: "authority.old"})

		input := authority.Authority{
			ID:   created.ID,
			Name: "authority.new",
		}

		updated, err := repo.Update(ctx, input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if updated.Name != "authority.new" {
			t.Errorf("expected name %q, got %q", "authority.new", updated.Name)
		}

		found, err := repo.FindByID(ctx, created.ID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if found.Name != "authority.new" {
			t.Errorf("renamed name not persisted: got %q", found.Name)
		}
	})

	t.Run("duplicate name", func(t *testing.T) {
		repo, ctx := newTestAuthorityMemoryRepository(t)
		target, _ := repo.Create(ctx, authority.Authority{Name: "authority.target"})
		other, _ := repo.Create(ctx, authority.Authority{Name: "authority.other"})

		input := authority.Authority{
			ID:   other.ID,
			Name: target.Name,
		}

		_, err := repo.Update(ctx, input)
		if err == nil {
			t.Fatal("expected error for duplicate name, got nil")
		}

		var appErr *domain.AppError
		if !errors.As(err, &appErr) {
			t.Fatalf("expected AppError, got %T", err)
		}

		if appErr.Cause != domain.ErrDuplicate {
			t.Errorf("expected ErrDuplicate, got %v", appErr.Cause)
		}
	})

	t.Run("index remains consistent", func(t *testing.T) {
		repo, ctx := newTestAuthorityMemoryRepository(t)
		created, _ := repo.Create(ctx, authority.Authority{Name: "authority.old"})

		input := authority.Authority{
			ID:   created.ID,
			Name: "authority.new",
		}

		_, err := repo.Update(ctx, input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if !repo.ExistByName("authority.new") {
			t.Error("expected new name to exist in nameIndex after update")
		}

		if repo.ExistByName("authority.old") {
			t.Error("expected old name to be removed from nameIndex after update")
		}
	})

	t.Run("duplicate update keeps original state", func(t *testing.T) {
		repo, ctx := newTestAuthorityMemoryRepository(t)

		target, err := repo.Create(
			ctx, authority.Authority{Name: "authority.target"},
		)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		other, err := repo.Create(
			ctx, authority.Authority{Name: "authority.other"},
		)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		_, err = repo.Update(ctx, authority.Authority{
			ID:   other.ID,
			Name: target.Name,
		})

		if err == nil {
			t.Fatal("expected duplicate error, got nil")
		}

		found, err := repo.FindByID(ctx, other.ID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if found.Name != "authority.other" {
			t.Errorf(
				"failed update modified entity: got %q",
				found.Name,
			)
		}
	})
}

func TestMemoryAuthorityRepository_DeleteByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo, ctx := newTestAuthorityMemoryRepository(t)
		created, _ := repo.Create(ctx, authority.Authority{Name: "authority.delete"})

		err := repo.DeleteByID(ctx, created.ID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("not found", func(t *testing.T) {
		repo, ctx := newTestAuthorityMemoryRepository(t)

		err := repo.DeleteByID(ctx, 999)
		if err == nil {
			t.Fatal("expected error for non-existent ID, got nil")
		}

		var appErr *domain.AppError
		if !errors.As(err, &appErr) {
			t.Fatalf("expected AppError, got %T", err)
		}

		if appErr.Cause != domain.ErrNotFound {
			t.Errorf("expected ErrNotFound, got %v", appErr.Cause)
		}
	})

	t.Run("removed from DB", func(t *testing.T) {
		repo, ctx := newTestAuthorityMemoryRepository(t)
		created, _ := repo.Create(ctx, authority.Authority{Name: "authority.delete"})

		err := repo.DeleteByID(ctx, created.ID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		_, err = repo.FindByID(ctx, created.ID)
		if err == nil {
			t.Fatal("expected error when finding deleted authority, got nil")
		}

		var appErr *domain.AppError
		if !errors.As(err, &appErr) {
			t.Fatalf("expected AppError, got %T", err)
		}

		if appErr.Cause != domain.ErrNotFound {
			t.Errorf("expected ErrNotFound, got %v", appErr.Cause)
		}
	})

	t.Run("removed from NameIndex", func(t *testing.T) {
		repo, ctx := newTestAuthorityMemoryRepository(t)
		created, _ := repo.Create(ctx, authority.Authority{Name: "authority.delete"})

		err := repo.DeleteByID(ctx, created.ID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if repo.ExistByName("authority.delete") {
			t.Error("expected name to be removed from NameIndex after delete")
		}
	})
}
