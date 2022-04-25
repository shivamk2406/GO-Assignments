<em>Topics Covered: Go Multithreading.</em>
 
Develop a multi-threaded go program where one thread reads the data from the database say, details of an Item from a mysql table. This thread builds an in-memory object, stores it in a collection. Simultaneously another thread should fetch already created Item objects from this collection and calculate the tax as per rules detailed in assignment#1 update the tax value in appropriate Item attribute and store it in a different collection. Finally print out the item details to console as detailed in assignment #1.
 
Implement such that the performance is optimal and thread race/deadlock is avoided.
 
<p><em>Real Life Scenario:</em><br>
Producer consumer mechanism.

**Key Points:**<br>
Please make sure your database is set up and you are able to access it before starting with implementation of this assignment.
Use Go’s multithreading support for implementation.
Proper validation / info messages should be thrown on the console wherever required.
Do appropriate error  handling wherever required.
Where ever required please write comments in the code to make it more understandable.
TDD methodology should be used


