package main

import (
   "fmt"
)

type Node struct {
    val int
    right *Node
    left *Node
}

func rightView( root *Node) {
    if (root == nil) { 
         return 
    }
    Q := make([]*Node, 2)
    Q[0] = root
    Q[1] = nil
    for (len(Q) != 0) {
        elem = Q[0]
        Q = Q[1:]
        prevelem:= elem
        If ( elem == nil ) {
            fmt.Printf("%d",prevelem.data)
            if (len(Q) != 0 ) {
                Q= append(Q,elem.right)
            }
            continue
        }
        if ( elem.left != nil) {
           Q= append(Q,elem.left)
        }
        if ( elem.right != nil) {
            Q= append(Q,elem.right)
        }
    } 
}

func main() {
  var view []int
  root := & Node {10,nil,nil}
  rightView(root)
}
