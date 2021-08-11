flutter build apk --split-per-abi &&^
flutter build web --web-renderer html --release &&^
xcopy /s /e /y build\web\ ..\docs