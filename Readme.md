

## CodeGen: A Command-Line Tool for Creating and Deleting Files
`codegen` is a command-line tool designed to create and delete files based on the provided name. It will create controller, service, repository, and entity classes for your Golang project.

## Installation

To install codegen, use the following command:
```bash
  sudo dpkg -i codegen-deb.deb  
```
    
## Usage
Create a File

To create a new file, use the create command followed by the desired name:

```
codegen create <name>
```

Delete a File

To delete an existing file, use the delete command followed by the name of the file:

```
codegen delete <name>
```
**Note :** The name parameter must be consistent when creating and deleting files.

Examples

Creating a file named example:
```
codegen create example
```
Deleting a file named example
```
codegen delete example
```

Hope that helps!