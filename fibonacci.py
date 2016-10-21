import datetime

def rec_fib(n):
    '''inefficient recursive function as defined, returns Fibonacci number'''
    if n > 1:
        return rec_fib(n-1) + rec_fib(n-2)
    return n

if __name__ == "__main__":
    start = datetime.datetime.now()
    for i in range(40):
        print(i, rec_fib(i))
    time = datetime.datetime.now() - start
    print(time)
