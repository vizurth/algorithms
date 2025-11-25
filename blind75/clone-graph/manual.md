### 133. Clone Graph
```go
 func dfs(node *Node, record map[int]*Node, visited map[*Node]bool) {
     if visited[node] || node == nil{
         return 
     }

     visited[node] = true
     newNode := &Node{}
     newNode.Val = node.Val

     record[node.Val] = newNode

     for _, child := range node.Neighbors {
         if !visited[child] {
             dfs(child, record, visited)
         }
         newNode.Neighbors = append(newNode.Neighbors, record[child.Val])
     }
 }


 func cloneGraph(node *Node) *Node {
     if node == nil {
         return nil
     }

     record := make(map[int]*Node)
     visited := make(map[*Node]bool)

     dfs(node, record, visited)

     return record[node.Val]
 }
```

```go
func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }

    record := make(map[int]*Node)
    record[node.Val] = &Node{Val: node.Val}

    bfs(node, record)

    return record[node.Val]
}

func bfs(startNode *Node, record map[int]*Node) {
    queue := []*Node{startNode}

    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]

        for _, neighbor := range node.Neighbors {
            if _, ok := record[neighbor.Val]; !ok {
                record[neighbor.Val] = &Node{Val: neighbor.Val}
                queue = append(queue, neighbor)
            }
            record[node.Val].Neighbors = append(record[node.Val].Neighbors, record[neighbor.Val])
        }
    }
}
```
### Описание алгоритма
Два способа решения через BFS и DFS