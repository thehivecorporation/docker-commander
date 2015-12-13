// Package container is Docker description
package container

// Container is a Docker entity
type Container struct {
	ID    int64  `form:"id" json:"id" binding:"required"`
	Title string `form:"title" json:"title" binding:"required"`
	State string `form:"state" json:"state" binding:"required"`
}
