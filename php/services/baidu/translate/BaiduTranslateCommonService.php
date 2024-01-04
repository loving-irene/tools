<?php

namespace services\baidu\translate;

use GuzzleHttp\Client;
use Monolog\Handler\StreamHandler;
use Monolog\Logger;

class BaiduTranslateCommonService
{
    public $appid;
    public $query;
    public $salt;
    public $secret;
    public $sign;

    public $param;
    public function init($appid,$secret)
    {
        $this->appid=$appid;
        $this->secret=$secret;
        $this->salt=random_int(1000000,9999999);
        return $this;
    }
    public function sign($query)
    {
        $this->query=$query;
        $str=$this->appid.$this->query.$this->salt.$this->secret;
        $this->sign=md5($str);
        return $this;
    }

    public function queryParam()
    {
       $this->param = [
           'q'=>$this->query,
           'from'=>'en',
           'to'=>'zh',
           'appid'=>$this->appid,
           'salt'=>$this->salt,
           'sign'=>$this->sign
       ];
       return $this;
    }
    public function translate()
    {
        $log=new Logger('info');
        $log->pushHandler(new StreamHandler(constant('root').'/info.log'));
        $log->info(__METHOD__,[$this->param]);
        $client=new Client(['base_uri'=>'https://fanyi-api.baidu.com/']);
        $response=$client->request('GET','api/trans/vip/translate',[
            'query'=>$this->param
        ]);

        if ($response->getStatusCode() === 200) {
            return $response->getBody()->getContents();
        }else{
            throw new \Exception('翻译请求错误'.$response->getReasonPhrase());
        }
    }
}
