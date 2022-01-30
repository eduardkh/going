
"""
second_function doesn't use anything from the main module
it is called with no parameters
"""


def second_function():
    print("This is the second function")


"""
third_function uses first_function from the main module
it is called with first_function as a parameter
"""


def third_function(first_function):
    print("This is the third function")
    first_function()
