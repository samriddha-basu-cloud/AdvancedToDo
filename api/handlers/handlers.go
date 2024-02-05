package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samriddha-basu-cloud/AdvancedToDo/db"
	"github.com/samriddha-basu-cloud/AdvancedToDo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateTask creates a new task.
func CreateTask(c *gin.Context) {
	var input struct {
		Title       string    `json:"title" binding:"required"`
		Description string    `json:"description" binding:"required"`
		DueDate     time.Time `json:"due_date" binding:"required"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := getUserIDFromContext(c)

	task := models.Task{
		Title:       input.Title,
		Description: input.Description,
		DueDate:     input.DueDate,
		Priority:    calculateTaskPriority(input.DueDate),
		Status:      "TODO",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := db.Collection.InsertOne(context.Background(), task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully"})
}

// CreateSubTask creates a new sub task.
func CreateSubTask(c *gin.Context) {
	var input struct {
		TaskID int `json:"task_id" binding:"required"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := getUserIDFromContext(c)

	// Check if task exists
	taskFilter := bson.M{"_id": input.TaskID}
	var existingTask models.Task
	err := db.Collection.FindOne(context.Background(), taskFilter).Decode(&existingTask)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create sub task"})
		}
		return
	}

	subTask := models.SubTask{
		TaskID:    input.TaskID,
		Status:    0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = db.Collection.InsertOne(context.Background(), subTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create sub task"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Sub task created successfully"})
}

// GetAllUserTasks retrieves all tasks for the user.
func GetAllUserTasks(c *gin.Context) {
	userID := getUserIDFromContext(c)

	var tasks []models.Task
	cursor, err := db.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode tasks"})
			return
		}
		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, tasks)
}

// GetAllUserSubTasks retrieves all sub tasks for the user.
func GetAllUserSubTasks(c *gin.Context) {
	userID := getUserIDFromContext(c)

	var subTasks []models.SubTask
	cursor, err := db.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch sub tasks"})
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var subTask models.SubTask
		if err := cursor.Decode(&subTask); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode sub tasks"})
			return
		}
		subTasks = append(subTasks, subTask)
	}

	c.JSON(http.StatusOK, subTasks)
}

// UpdateTask updates a task.
func UpdateTask(c *gin.Context) {
	taskID := c.Param("id")

	var update struct {
		DueDate time.Time `json:"due_date"`
		Status  string    `json:"status"`
	}

	if err := c.BindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := getUserIDFromContext(c)

	// Update task logic here

	c.Status(http.StatusOK)
}

// UpdateSubTask updates a sub task.
func UpdateSubTask(c *gin.Context) {
	subTaskID := c.Param("id")

	var update struct {
		Status int `json:"status"`
	}

	if err := c.BindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := getUserIDFromContext(c)

	// Update sub task logic here

	c.Status(http.StatusOK)
}

// DeleteTask soft deletes a task.
func DeleteTask(c *gin.Context) {
	taskID := c.Param("id")

	userID := getUserIDFromContext(c)

	// Soft delete task logic here

	c.Status(http.StatusOK)
}

// DeleteSubTask soft deletes a sub task.
func DeleteSubTask(c *gin.Context) {
	subTaskID := c.Param("id")

	userID := getUserIDFromContext(c)

	// Soft delete sub task logic here

	c.Status(http.StatusOK)
}

func calculateTaskPriority(dueDate time.Time) int {
	today := time.Now().Truncate(24 * time.Hour)
	diff := dueDate.Sub(today).Hours() / 24

	if diff == 0 {
		return 0
	} else if diff <= 2 {
		return 1
	} else if diff <= 4 {
		return 2
	} else {
		return 3
	}
}

func getUserIDFromContext(c *gin.Context) int {
	// Retrieve user ID from JWT token in context or session
	// Example: userID := c.GetInt("user_id")
	// Return userID or handle error accordingly
	return 1 // Placeholder value, replace with actual user ID retrieval logic
}
