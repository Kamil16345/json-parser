# JSON Parser

## Goals

This project is a command-line interface (CLI) application that can parse JSON data and convert it into a structured Go data structure. Additionally, the application is able to generate JSON data from a given Go structure. The application handles proper error handling, structure validation, and IO operations.

## Project assumptions

1. JSON Parsing
    - The application should read JSON data from a file or accept JSON input from the user.
    - It should parse the JSON data and convert it into a structured Go data structure.
2. Go Structure Parsing
    - The application should accept a Go structure as a parameter and generate JSON data from it.
    - It should validate the provided Go structure and ensure it matches the expected format.
3. Error Handling
    - The application should handle JSON parsing errors and Go structure validation errors gracefully.
    - It should provide informative error messages to guide the user in resolving the issues.
4. Structure Validation
    - The application should validate the parsed JSON structure against the provided Go structure.
    - It should ensure that the JSON data matches the expected structure and data types defined in the Go structure.
    - Likewise, it should validate the Go structure against the provided JSON data.
5. IO Operations
    - The application should support reading JSON data from a file and writing JSON data to a file.
    - It should provide options to specify the input source (file path or user input) and the output destination (file path or console).
6. Go Structure Generation
    - The application should dynamically generate a Go structure based on the parsed JSON data.
    - It should generate a Go structure that accurately represents the JSON data, including nested objects, arrays, and primitive types.
7. Bi-Directional Conversion
    - The application should be capable of both JSON-to-Go-structure parsing and Go-structure-to-JSON generation.
    - It should provide separate commands or options to perform each conversion.
8. Command-Line Interface
    - The application should have a clear and intuitive CLI with well-defined commands and options.
    - The CLI should display instructions, usage guidelines, and error messages when necessary.
9. Error Reporting
    - The application should provide detailed error reporting in case of any issues during JSON parsing, structure validation, or Go structure generation.
    - Error messages should be informative and guide the user on how to resolve the errors.
