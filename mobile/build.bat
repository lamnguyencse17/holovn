flutter build apk --split-per-abi &&^
flutter build web --web-renderer html --release &&^
xcopy /s /e /y build\web\ ..\docs &&^
move /f build\app\outputs\apk\release\app-arm64-v8a-release.apk build\app\outputs\apk\release\holovn-arm64-v8a-release.apk &&^
move /f build\app\outputs\apk\release\app-armeabi-v7a-release.apk build\app\outputs\apk\release\holovn-armeabi-v7a-release.apk &&^
move /f build\app\outputs\apk\release\app-x86_64-release.apk build\app\outputs\apk\release\holovn-x86_64-release.apk