<?php

use Symfony\Component\DomCrawler\Crawler;

require __DIR__ . '/../../vendor/autoload.php';

$html = <<<'HTML'
<!DOCTYPE html>
<html>
    <body>
        <p class="message">Hello World!</p>
        <p>Hello Crawler!</p>
    </body>
</html>
HTML;

$crawler = new Crawler($html);

foreach ($crawler as $domElement) {
  var_dump($domElement->nodeName);
}

