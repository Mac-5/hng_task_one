# Number Classification API Documentation
## Overview 
This project provides a simple public API built using Go. The Number Classification API allows users to input a number and receive information about its mathematical properties, including:

- Prime: Whether the number is prime.
- Perfect: Whether the number is a perfect number.
- Armstrong: Whether the number is an Armstrong number.
- Even/Odd: Whether the number is even or odd.
- Digit Sum: The sum of the number's digits.
- Fun Fact: An interesting fact about the number. 

## Table of Contents
- [Prerequisites](#PREREQUISITES)
- [Setup](#SETUP)
-  [Endpoints](#APi%Documentation)
- [Example Usage](#Example%Usage)
- [License](#license)
 
 ## PREREQUISITES
 
 Before running this project, ensure that the following software is installed on your local machine:
  - Go (version 1.18 or higher)
 
## SETUP

### Running the Project Locally Follow these steps to set up and run the project:

1. **Clone the repository:** 
 git clone [https://github.com/Mac-5/hng_task_one.git](https://github.com/Mac-5/hng_task_one.git)
 
 2. **Navigate into the project directory:**
```
		cd hng_task_one  
```

3. **Run the Go server:**
```
		go run main.go
```
## API Documentation

### Endpoint URL

-   URL: `GET  `

### Request Format

-   Method: `GET`
-   URL: `https://`
-   Headers:
    -   `Content-Type: application/json`

### Response Format

The API responds with a JSON object containing the following fields:

- number (integer): The input number.
- is_prime (boolean): Indicates if the number is prime.
- is_perfect (boolean): Indicates if the number is a perfect number.
- properties (array of strings): A list of properties the number possesses (e.g., "prime", "perfect", "armstrong", "even", "odd").
- digit_sum (integer): The sum of the number's digits.
- fun_fact (string): An interesting fact about the number.
```json
{
    "number": 371,
    "is_prime": false,
    "is_perfect": false,
    "properties": ["armstrong", "odd"],
    "digit_sum": 11,  
    "fun_fact": "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371"
}
```
## Example Usage

To interact with the API, send a `GET` request to the endpoint `/`.

Use `curl` to test the API:
```sh
curl -X GET https://hng-task-zero.onrender.com/
```

Expected JSON response:
```json
{
    "number": 371,
    "is_prime": false,
    "is_perfect": false,
    "properties": ["armstrong", "odd"],
    "digit_sum": 11,
    "fun_fact": "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371"
}

```

## Backlinks

[hng-hire-golang-devs](https://hng.tech/hire/golang-developers)

## License

-
