package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"restapi/todo"
	"time"

	"github.com/gorilla/mux"
)

type HTTPHandlers struct {
	todoList *todo.List
}

func NewHTTPHandlers(todoList *todo.List) *HTTPHandlers {
	return &HTTPHandlers{
		todoList: todoList,
	}
}

//REST API config for HandlerCreateTask
/*1 Request
pattern /tasks
method/ POST
add. info: JSON in http-request body

2 Response

2.1 succeed
-status code: 201 (created)
-response body: JSON represens created task

2.2 failed
-status code: 400, 409, 500... (error)
 -response body: JSON+error + time
*/
func (h *HTTPHandlers) HandlerCreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDTO TaskDTO
	//validation of json parsing
	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		//struct error-handler
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	//validation that Title and Description are not empty
	if err := taskDTO.ValidateFromCreate(); err != nil {
		//struct error-handler
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	//creation of new task
	todoTask := todo.NewTask(taskDTO.Title, taskDTO.Description)

	//add task to list
	if err := h.todoList.AddTask(todoTask); err != nil {
		//struct error-handler
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		if errors.Is(err, todo.ErrTaskAlreadyExists) {
			http.Error(w, errDTO.ToString(), http.StatusConflict)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}
	b, err := json.MarshalIndent(todoTask, "", "	")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Fail to write http-response:", err)
		return
	}

}

//REST API config for HandlerGetTask
/*1 Request
pattern /tasks/{title}
method/ GET
add. info: see pattern

2 Response

2.1 succeed
-status code: 200 (OK)
-response body: JSON represens data from task

2.2 failed
-status code: 400, 409, 500... (error)
 -response body: JSON+error + time
*/
func (h *HTTPHandlers) HandlerGetTask(w http.ResponseWriter, r *http.Request) {
	//title, ok := mux.Vars(r)["title"] //no need to check for ok, since
	//router.Path("/tasks/{title}")....
	title := mux.Vars(r)["title"]

	//get task
	task, err := h.todoList.GetTask(title)
	if err != nil {
		//struct error-handler
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		if errors.Is(err, todo.ErrTaskNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}
	b, err := json.MarshalIndent(task, "", "	")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Fail to write http-response:", err)
		return
	}

}

//REST API config for HandlerGetAllTasks
/*1 Request
pattern/tasks
method/ GET
add. info: -

2 Response

2.1 succeed
-status code: 200 (OK)
-response body: JSON represens data from tasks

2.2 failed
-status code: 400, 500... (error)
 -response body: JSON+error + time
*/
func (h *HTTPHandlers) HandlerGetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.todoList.ListTasks()
	b, err := json.MarshalIndent(tasks, "", "	")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Fail to write http-response:", err)
		return
	}
}

//REST API config for HandlerGetAllUncomletedTasks
/*1 Request
pattern/tasks?uncompetes=true
method/ GET
add. info: querry parameters (for filter)

2 Response

2.1 succeed
-status code: 200 (OK)
-response body: JSON represens data from tasks

2.2 failed
-status code: 400, 500... (error)
 -response body: JSON+error + time
*/
func (h *HTTPHandlers) HandlerGetAllUncomletedTasks(w http.ResponseWriter, r *http.Request) {
	uncomletedTasks := h.todoList.ListUncompletedTasks()

	b, err := json.MarshalIndent(uncomletedTasks, "", "	")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Fail to write http-response:", err)
		return
	}
}

//REST API config for HandlerCompleteTask
/*1 Request
pattern /tasks/{title}
method/ PATCH
add. info: pattern + JSON in http-request body (field Complete to chahge)

2 Response
2.1 succeed
-status code: 200 (OK)
-response body: JSON represens data from changed tasks

2.2 failed
-status code: 400, 409, 500... (error)
 -response body: JSON+error + time
*/
func (h *HTTPHandlers) HandlerCompleteTask(w http.ResponseWriter, r *http.Request) {
	var comleteTaskDTO CompleteTaskDTO
	//validation of json parsing
	if err := json.NewDecoder(r.Body).Decode(&comleteTaskDTO); err != nil {
		//struct error-handler
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	//read title
	title := mux.Vars(r)["title"]

	var (
		changedTask todo.Task
		err         error
	)

	if comleteTaskDTO.Complete {
		changedTask, err = h.todoList.CompleteTask(title)
	} else {
		changedTask, err = h.todoList.UncompleteTask(title)
	}

	if err != nil {
		//struct error-handler
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		if errors.Is(err, todo.ErrTaskNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}

	/*if comleteTaskDTO.Complete {
		changedTask, err := h.todoList.CompleteTask(title)
		if err != nil {
			//struct error-handler
			errDTO := ErrorDTO{
				Message: err.Error(),
				Time:    time.Now(),
			}
			if errors.Is(err, todo.ErrTaskNotFound) {
				http.Error(w, errDTO.ToString(), http.StatusNotFound)
			} else {
				http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
			}
			return
		} else {
			changedTask, err := h.todoList.UncompleteTask(title)
			if err != nil {
				//struct error-handler
				errDTO := ErrorDTO{
					Message: err.Error(),
					Time:    time.Now(),
				}
				if errors.Is(err, todo.ErrTaskNotFound) {
					http.Error(w, errDTO.ToString(), http.StatusNotFound)
				} else {
					http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
				}
				return
			}
		}
	}*/

	b, err := json.MarshalIndent(changedTask, "", "	")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Fail to write http-response:", err)
		return
	}

}

//REST API config for HandlerDeleteTask
/*1 Request
pattern /tasks/{title}
method/ DELETE
add. info: pattern

2 Response
2.1 succeed
-status code: 204 (No Content)
-response body: -

2.2 failed
-status code: 400, 404, 500... (error)
 -response body: JSON+error + time
*/
func (h *HTTPHandlers) HandlerDeleteTask(w http.ResponseWriter, r *http.Request) {
	//read title
	title := mux.Vars(r)["title"]
	if err := h.todoList.DeleteTask(title); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		if errors.Is(err, todo.ErrTaskNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
