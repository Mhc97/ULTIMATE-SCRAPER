<?php
$books = [];
$url = "https://books.toscrape.com/catalogue/page-1.html";
for ($i=0; $i<5; $i++){
    $dom = new DOMDocument(); 
    @$dom->loadHTML(file_get_contents($url));
    $x = new DOMXPath($dom);
    $items = $x->query("//article[@class='product_pod']");
    foreach ($items as $a){
        $t = $x->query(".//h3/a", $a)->item(0);
        $books[] = ['title' => $t ? $t->getAttribute("title") : ''];
    }
    $n = $x->query("//li[@class='next']/a")->item(0);
    if (!$n) break;
    $url =  "https://books.toscrape.com/catalogue/" . $n->getAttribute('href');
}
if (!is_dir('data')) mkdir('data');
file_put_contents('data/PHP.json', json_encode($books));
echo count($books)
?>