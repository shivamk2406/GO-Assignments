<h1> Dependency-Graph </h1>
</em><strong>Topics Covered: Data Structures In Go .</strong></em><br>

Design a Data Structure that represents a dependency graph.

Dependency Graph is an acyclic multi root directional graph with the exception of a root node, which has no parents.

<p><em>Real Life Scenario:</em><br>
Family Tree</p>

<p><em>Terminology used:</em><br>
Parent: For edge A->B, A is a parent of B. There may be multiple parents for a child.<br>
Child: For edge A->B, B is a child of A. There may be multiple children of a parent.<br>
Ancestor: parent or grand-parent or grand-grand-parent and so on<br>
Descendant: child or grand-child or grand-grand-child and so on<br></p>

Basically the data structure should allow you to store the parent child relationship and this can go to the nth level.

<p> <em>Design:</em><br>
The node information, which we will store, is:<br>
Node Id --- This has to be unique.<br>
Node Name. Need not be distinct.<br>
Additional Information --- In the form of a key value pairs and this can be different for each node.<br></p>

<p><em>Operations:</em><br>

Get the immediate parents of a node, passing the node id as input parameter.<br>
Get the immediate children of a node, passing the node id as input parameter.<br>
Get the ancestors of a node, passing the node id as input parameter.<br>
Get the descendants of a node, passing the node id as input parameter.<br>
Delete dependency from a tree, passing parent node id and child node id.<br>
Delete a node from a tree, passing node id as input parameter. This should delete all the dependencies of the node.<br>
Add a new dependency to a tree, passing parent node id and child node id. This should check for cyclic dependencies.<br>
Add a new node to tree. This node will have no parents and children. Dependency will be established by calling the 7 number API.<br>

<p><strong>Key Points:</strong><br>
Use go data structures to implement Family Tree.<br>
Proper validation / info messages should be thrown on console.</br>
Do appropriate error  handling wherever required.</br>
Where ever required please write comments in the code to make it more understandable.</br>
TDD methodology should be used<br></p>





