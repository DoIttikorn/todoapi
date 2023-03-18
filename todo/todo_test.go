package todo

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TestSleepNotProcess(t *testing.T) {
	handler := NewTodoHandler(&gorm.DB{})

	w := httptest.NewRecorder()
	payload := bytes.NewBufferString(`{"text": "sleep"}`)
	req := httptest.NewRequest("POST", "/todos", payload)
	req.Header.Set("TransactionID", "testID123")

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	handler.NewTask(c)

	want := `{"error":"not allowed"}`

	if w.Body.String() != want {
		t.Errorf("NewTask() = %v, want %v", w.Body.String(), want)
	}

}
