#!/usr/bin/env python3

import requests

def get_joke():
    url = "https://icanhazdadjoke.com/"
    headers = {"Accept": "application/json"}
    response = requests.get(url, headers=headers, timeout=5)
    joke = response.json()["joke"]
    return joke


if __name__ == '__main__':
    print(get_joke())


