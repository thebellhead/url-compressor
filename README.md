# Проект: Сервис коротких ссылок

В рамках данных задач стоит амбициозный и креативный проект по созданию сервиса коротких ссылок, который удобно вписывается в современный веб-ландшафт. Наш сервис будет принимать на вход стандартные REST запросы, содержащие оригинальные URL-адреса, и выдавать в ответ компактные, укороченные версии с использованием домена `localhost`. Пользователи, перейдя по такой короткой ссылке, будут автоматически перенаправлены на изначальный, полный адрес ресурса.

В качестве образца можно рассматривать сервисы типа https://surl.li/ru, которые демонстрируют функционал и потенциал сокращения ссылок.

В результате проектных работ ожидается:
1. Подробное архитектурное описание с тщательным анализом каждого принятого решения. Здесь будут освещены такие аспекты, как причины выделения функциональности в отдельный микросервис, выбор способа коммуникации – Kafka/GRPC, логика за выбором определённого типа базы данных, и другие ключевые моменты.
2. Полноценная реализация сервиса, отвечающая всем поставленным требованиям и стандартам качества. Не забудьте написать тесты для вашего проекта.
3. Docker compose файл, содержащий все необходимые настройки для быстрого и безболезненного запуска сервиса в любой среде.
4. Документация интерфейса сервиса, включающая в себя спецификацию REST запроса для генерации короткой ссылки и прочие важные детали взаимодействия с сервисом.
