<?php

namespace services\baidu\translate;

class BaiduTranslateFactory
{

    public static function createTranslate($type)
    {
        switch ($type) {
            case 'common':
                return new BaiduTranslateCommonService();
            case 'domain':
                return new BaiduTranslateDomainService();
        }
    }
}
