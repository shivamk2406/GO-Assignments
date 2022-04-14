<h1> Dependency-Graph </h1>
</em><strong>Topics Covered: Data Structures In Go .</strong></em>

Design a Data Structure that represents a dependency graph.

Dependency Graph is an acyclic multi root directional graph with the exception of a root node, which has no parents.

*Real Life Scenario:*
Family Tree

*Terminology used:*
Parent: For edge A->B, A is a parent of B. There may be multiple parents for a child.
Child: For edge A->B, B is a child of A. There may be multiple children of a parent.
Ancestor: parent or grand-parent or grand-grand-parent and so on
Descendant: child or grand-child or grand-grand-child and so on

Basically the data structure should allow you to store the parent child relationship and this can go to the nth level.

*Design:*
The node information, which we will store, is:
Node Id --- This has to be unique.
Node Name. Need not be distinct.
Additional Information --- In the form of a key value pairs and this can be different for each node.

*Operations:*

Get the immediate parents of a node, passing the node id as input parameter.
Get the immediate children of a node, passing the node id as input parameter.
Get the ancestors of a node, passing the node id as input parameter.
Get the descendants of a node, passing the node id as input parameter.
Delete dependency from a tree, passing parent node id and child node id.
Delete a node from a tree, passing node id as input parameter. This should delete all the dependencies of the node.
Add a new dependency to a tree, passing parent node id and child node id. This should check for cyclic dependencies.
Add a new node to tree. This node will have no parents and children. Dependency will be established by calling the 7 number API.

**Key Points:**
Use go data structures to implement Family Tree.
Proper validation / info messages should be thrown on console.
Do appropriate error  handling wherever required.
Where ever required please write comments in the code to make it more understandable.
TDD methodology should be used





