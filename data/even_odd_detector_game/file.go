def read_number_from_file(file_path):
    with open(file_path, 'r') as file:
        number = int(file.read().strip())
    return number

def is_even(number):
    return number % 2 == 0

def main():
    # Path to the file where the number is stored
    file_path = 'data/number.txt'
    
    # Read the number from the file
    number = read_number_from_file(file_path)
    
    # Check if the number is even or odd
    if is_even(number):
        print(f"The generated number {number} is even.")
    else:
        print(f"The generated number {number} is odd.")

if __name__ == "__main__":
    main()
