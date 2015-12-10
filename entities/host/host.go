/*
 * A host entity. A host is a remote machine that has containers. Has a 1 to
 * many relationship with containers
 */

package host

import (
  // "container/list"
  // "github.com/sayden/docker-commander/entities/container"
)

type Host struct {
  Id int64
  Title string
  Ip string
}
