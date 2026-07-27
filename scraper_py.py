import requests, json, os
from bs4 import BeautifulSoup

books, url = [], "https://books.toscrape.com/catalogue/page-1.html"
for _ in range(5):
    soup = BeautifulSoup(requests.get(url).text, 'html.parser')
    items = soup.select('article.product_pod')
    for a in items:
        books.append({"title": a.h3.a['title']})
    n = soup.select_one('li.next a')
    if not n: 
        break
    url = "https://books.toscrape.com/catalogue/" + n['href']
os.makedirs("data", exist_ok=True)
with open("data/Python.json","w") as f:
    json.dump(books, f)
print(len(books))