# Telegram webapp flea-market

[![Release](https://github.com/konstantin-kukharev/flea-market/actions/workflows/release.yml/badge.svg)](https://github.com/konstantin-kukharev/flea-market/actions/workflows/release.yml)
[![Lint & test](https://github.com/konstantin-kukharev/flea-market/actions/workflows/lint.yml/badge.svg)](https://github.com/konstantin-kukharev/flea-market/actions/workflows/lint.yml)
[![CI](https://github.com/konstantin-kukharev/flea-market/actions/workflows/ci.yml/badge.svg)](https://github.com/konstantin-kukharev/flea-market/actions/workflows/ci.yml)
___
для корректного деплоя необходимо создать секреты в репозитории
```
- secrets.CI_TOKEN
- secrets.TELEGRAM_TOKEN # Tg bot token
- secrets.TELEGRAM_TO # Sending deploy message
```
для сборки тестового окружения используется [go task](https://taskfile.dev)
___
Соседские чаты изначально создавались как некий канал для общения соседей и предоставления важной информации от домоуправления (ЖЭКа, управляющая компания). Сегодня люди редко ходят на общедомовые собрания и практически не читают объявления у подъезда, однако они живут в квартирах и имеют право влиять на пространство вокруг себя.
    
Также до них нужно как-то доносить важную информацию о парковках, смене тарифа, изменении УК или строительстве детской площадки во дворе. Поквартирные обходы – не вариант: этим просто никто не будет заниматься. Гораздо проще и эффективнее использовать чат, в котором за пару минут можно сообщить важные сведения всем своим соседям.

Соседские чаты заменили скамейки во дворе. До их появления все новости округи можно было узнать у старушек возле подъезда: сейчас же эти бабушки перекочевали в смартфон. Но в чаты пишут не только они. Практически каждый человек, столкнувшийся с чем-то интересным или ужасным в районе, стремится выложить это на всеобщее обозрение. Новости обрастают комментариями, и в итоге соседский чат превращается в настоящую черную дыру. В мессенджере можно прочитать о:
- ДТП (заботливые соседи даже скинут фото искореженной машины);
- краже велосипеда у подъезда (у кого-то моментально найдутся записи с видеокамер);
- утере ключей ребенком (дадут ссылку на пост того, кто их нашел и даже разместил в чат часом ранее);
- поисках потерявшейся собаки (кто-нибудь обязательно вспомнит, что видел ее у магазина в двух кварталах от дома).

С одной стороны чаты полезны, но с другой – отнимают массу времени. Действительно важные сообщения (к примеру, объявления управляющей компании) часто теряются за ворохом других малоценных постов.

Данное приложение решает вопрос с организацией информации о продаже/покупке/... вещей, как правило вся информация такого характера организуется в отдельном чате под названием "Барахолка"

Функционал чата:
- не дает информации об актуальности объявлений
- не делит сообщения на категории (ремонт/садоводство/...)
- нет возможности отфильтровать в чате отклики на размещенное объявление