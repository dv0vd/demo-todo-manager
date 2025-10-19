package tests

import (
	"demo-todo-manager/internal/dto"
	responses "demo-todo-manager/internal/http/responses/note"
	testutils "demo-todo-manager/internal/tests"
	"strconv"
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestNoteResponse(t *testing.T) {
	id, _ := strconv.ParseUint(testutils.GetRandomInt(100, 1000), 10, 64)
	note := dto.NoteDTO{
		ID:          id,
		Title:       faker.Word(),
		Description: faker.Sentence(),
	}

	testutils.CheckResult(
		t,
		"Note response",
		responses.NoteResponse(note),
		responses.NoteResponseStruct{
			Success: true,
			Data: responses.NoteData{
				ID:          note.ID,
				Title:       note.Title,
				Description: note.Description,
			},
		},
	)
}

func TestNotesResponse(t *testing.T) {
	notesCount, _ := strconv.ParseUint(testutils.GetRandomInt(3, 10), 10, 32)
	notesDTO := []dto.NoteDTO{}
	notes := []responses.Note{}

	for i := uint64(0); i <= notesCount; i++ {
		id, _ := strconv.ParseUint(testutils.GetRandomInt(100, 1000), 10, 64)
		title := faker.Word()
		description := faker.Sentence()

		notesDTO = append(notesDTO, dto.NoteDTO{
			ID:          id,
			Title:       title,
			Description: description,
		})
		notes = append(notes, responses.Note{
			ID:          id,
			Title:       title,
			Description: description,
		})
	}

	testutils.CheckResult(
		t,
		"Notes response",
		responses.NotesResponse(notesDTO),
		responses.NotesResponseStruct{
			Success: true,
			Data: responses.NotesData{
				Notes: notes,
			},
		},
	)
}
