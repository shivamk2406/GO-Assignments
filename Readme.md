Write a menu driven command line go program that provides the following menu options:
1.  Add User details.
2.  Display User details.
3.  Delete User details
4.  Save User details.
5.  Exit

The option (1) should ask for the following user details. All the following details are mandatory and the program should perform the required validations (like blank data, integer value for age, roll number etc). Roll Number is a key to identify the uniqueness among the students.
1.  Full Name
2.  Age
3.  Address
4.  Roll Number
5.  Set of courses he is interested to enroll. There are a total of 6 courses (Course A, B, C, D, E and F). It is mandatory for each student to choose 4 out of 6 courses.

Once the validations are passed the user details should be added to an in memory data structure. The data structure should always keep the records sorted in ascending order. By default the records should be sorted on full name. If the name is the same for two students then sorting should be based on the roll number.

The option (2) should display the user details in the following format. Also the user should be provided with an option to sort the results (either in ascending or descending order) based on name, roll number, age, address.

----------------------------------------------------------------------------------------------------------
Name    Roll Number                       Age                      Address                        Courses
---------------------------------------------------------------------------------------------------------- 
A            43                           1                        22 A, GGn                    A, C, D, E

The option (3) should ask for roll number and delete the student details corresponding to that roll number. Throw a proper user friendly message in case the roll number entered by the user does not exist.


The option (4) should save the in memory details of all the users to a disk. Use go’s serialization capabilities to serialize the in memory data to disk. If the user terminates the program after choosing this option the user’s data should be saved to disk and next time the user runs the program the in-memory collection should be pre populated with data already stored on the disk. 

The option (5) should terminate the program but before termination it should ask the user if he wants to save his latest changes (additions, deletions of users) to disk.

Key Points:
1.  Use Go’s serialization mechanism to save user details to disk.
2.  Use Go proper pkgs/libs  for sorting.
3.  Coding conventions should be followed.
4.  Proper validation / info messages should be thrown on console.
5.  Student Info, course info, serialization code and command line menu code should be encapsulated in    separate independent java classes.
6.  Where ever required please write comments in the code to make it more understandable.
7.  TDD methodology should be used
 

