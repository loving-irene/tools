<?php

require_once './vendor/autoload.php';

use League\Csv\Reader;

function myAutoloader($className) {
    // 将命名空间中的反斜杠转换为目录分隔符
    $className = str_replace('\\', DIRECTORY_SEPARATOR, $className);

    // 定义类文件的目录
    $classFile = __DIR__ . DIRECTORY_SEPARATOR . $className . '.php';

    // 检查类文件是否存在，如果存在则包含它
    if (file_exists($classFile)) {
        require_once $classFile;
    }
}

// 注册自动加载函数
spl_autoload_register('myAutoloader');

define('root',__DIR__);

$reader=Reader::createFromPath(__DIR__.'/storage/words.csv','r+');
$records=$reader->getRecords();
$writer=\League\Csv\Writer::createFromPath(__DIR__.'/storage/words_translate.csv','w+');

$baidu=new \services\baidu\translate\BaiduTranslateCommonService('20200428000432904','7Tw3AOdMtCgMpIScsTaq');
$baidu=\services\baidu\translate\BaiduTranslateFactory::createTranslate('domain');
$data=[];
foreach ($records as $record) {
    //TODO 将配置拆分到env中
    $res=$baidu->init('','')->sign($record[0])->queryParam()->translate();
    $decode=json_decode($res,true);
    if (array_key_exists('error_code', $decode)) {
        //出错
    }else{
        $data[]=[$decode['trans_result'][0]['src'],$decode['trans_result'][0]['dst']];
    }
}
$writer->insertAll($data);
