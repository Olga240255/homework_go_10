# Единый порядок домашней работы

Выполняйте разделы строго по порядку. В каждом разделе пять задач. Сначала прочитайте рассуждение, затем найдите `TODO(XX.Y)`, внесите одно изменение и запустите короткую проверку.

| № | Раздел | Документ | Проверка |
|---|---|---|---|
| 01 | JSON и модели | `01_json_models.md` | `go test ./internal/note -run JSON` |
| 02 | Store | `02_store.md` | `make test-store` |
| 03 | Router | `03_router.md` | `make test-http` |
| 04 | GET handlers | `04_get_handlers.md` | `make test-http` |
| 05 | POST handler | `05_post_handler.md` | `make test-http` |
| 06 | PATCH и DELETE | `06_patch_delete.md` | `make test-http` |
| 07 | Ошибки и статусы | `07_errors_statuses.md` | `make test` |
| 08 | HTTP-тесты и ручная проверка | `08_tests_manual.md` | `make ci` |

Стартовая версия обязана проходить `make check-start`, но функциональные тесты будут красными. Это ожидаемо: тесты показывают следующий невыполненный контракт.
