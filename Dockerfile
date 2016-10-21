FROM index.alauda.cn/library/python:3

COPY fibonacci.py /

CMD ["python", "/fibonacci.py"]