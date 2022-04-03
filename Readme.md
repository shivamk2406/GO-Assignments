Write a go program that accepts details (item name, item type, item prize) of different items from
Command line and outputs the item name, item prize, sales tax liability per item, final prize (sales tax + item prize) to the console. The input should be accepted with following command line options:
-name <first item name>
-price <price of first item>
-quantity <quantity of first item>
-type <type of first item>

The following functionalities/behavior is expected:
1. All options other than item name can come in any order i.e. after -name you can have -price, -type option. Item type is a mandatory option. 
2. The system should provide functionality to add more than one items i.e. Once the details of the first item is entered it should display a message saying:
    Do you want to enter details of any other item (y/n):
  	Appropriate behavior is required to be implemented as per user input for this question.
3. Make use of go’s object oriented capabilities for implementing this business logic.
4. Item type can have 3 possible values: raw, manufactured and imported.
Tax rules for the 3 types are as follows:
    1. raw: 12.5% of the item cost
    2. manufactured: 12.5% of the item cost + 2% of (item cost + 12.5% of the item cost)
    3. imported: 10% import duty on item cost + a surcharge (surcharge is: Rs. 5 if the final cost after applying tax & import duty is up to Rs. 100, Rs. 10 if the cost exceeds 100 and up to 200 and 5% of the final cost if it exceeds 200).

Key Points:
 
1. Use Go’s I/O capabilities to accept input from users.
2. Use Go’s String functionalities to parse input strings.
3. Coding conventions should be followed.
4. Proper validation / info messages should be thrown on console.
5. Do appropriate error handling wherever required.
6. Where ever required please write comments in the code to make it more understandable.
7. TDD methodology should be used

