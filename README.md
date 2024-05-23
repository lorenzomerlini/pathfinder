# Pathfinder Tool

Pathfinder Tool is a command-line utility written in Go that allows users to search for specific files or files with a specific extension within their file system. It provides a simple and efficient way to navigate through directories and locate desired files using concurrent goroutines, enabling multiprocess search capabilities.
Pathfinder Tool is designed to be fast, lightweight, and user-friendly, making it ideal for both casual and professional users who need to quickly locate files within their file system.


## Basic features 

- Searches for specific files in filesystem
- Uses goroutines to perform searches in parallel
- Allows multiple searches in succession, without restarting the program
- Clear terminal between searches for better readability
- Print the results or save them in a .txt file

## Features added

- **Configurable maximum goroutines**
- **Search entire filesystem**
- **Search by file extension**

## Future updates

- **Parallel search with different priorities**
- **Detailed logging**

## Requirements

Go 1.16 or higher. 

## Author 

Lorenzo Merlini
