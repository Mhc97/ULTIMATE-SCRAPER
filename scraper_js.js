const axios = require('axios');
const cheerio = require('cheerio');
const fs = require('fs');

(async () => {
    let b = [], u = "https://books.toscrape.com/catalogue/page-1.html";
    for (let i = 0; i < 5; i++){
      
        const $ = cheerio.load((await axios.get(u)).data);
        const items = $('article.product_pod');
        items.each((_, e) =>{
            b.push({ title: $(e).find('h3 a').attr('title')});
        });
    const n = $('li.next a');
    if (!n.length) break;
    u = "https://books.toscrape.com/catalogue/" + n.attr('href');
    }
    fs.mkdirSync('data', { recursive: true});
    fs.writeFileSync('data/Node.json', JSON.stringify(b));
    console.log(b.length);
})();