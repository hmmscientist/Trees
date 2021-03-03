package main
import (
    "fmt"
)
type Node struct {
   val int
   right *Node
   left *Node
}

func LeftView(root *Node) {

    if ( root == nil) return ;

    Q := make([]*Node,2)
    Q[0] = root
    Q[1] = nil
    for  (len(Q) != 0) {
         elem = Q[0]
         elem = Q[1:]
          if elem != nil && Q[0] == nil {
              fmt.Printf ("%d",elem.val)
          } else if elem == nil {
              continue
          }
          if (elem.left != nil ) {
              Q= append(Q,elem.left)
          }
          if (elem.right != nil ) {
              Q= append(Q,elem.right)
          }
          if (len(Q) != 0 && elem == nil ) {
              Q = append(Q,nil)
          }
    }
}

func main () {
    root := &Node {40,nil,nil}
    LeftView(root)
}
