from mod import second_function, third_function


def main():
    """
    first function calls second function (imported from mod.py)
    """
    def first_function():
        print("This is the first function")
        second_function()

    first_function()
    """
    third function (imported from mod.py) calls first function
    gets it as a parameter
    """
    third_function(first_function)


if __name__ == "__main__":
    main()
