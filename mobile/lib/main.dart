import 'package:flutter/material.dart';
import 'package:holovn_mobile/router/route_parser.dart';
import 'package:holovn_mobile/router/router.dart';
import 'package:get/get.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:flutter/foundation.dart';

Future main() async {
  try {
    if (!kReleaseMode){
      await dotenv.load(fileName: "assets/.env");
    }
  } catch (err) {
    print(err);
  }

  runApp(HolovnApp());
}

class HolovnApp extends StatefulWidget {
  @override
  State<StatefulWidget> createState() => _HolovnAppState();
}

class _HolovnAppState extends State<HolovnApp> {
  final routerDelegate = Get.put(AppRouter(), tag: "router");
  void initState() {
    super.initState();
  }

  RouteParser _routeInformationParser = RouteParser();

  @override
  Widget build(BuildContext context) {
    return MaterialApp.router(
      title: "Holovn - A Vietnamese Hololive Fan App",
      theme: ThemeData.dark().copyWith(
          brightness: Brightness.dark,
          // primarySwatch: Colors.pink,
          primaryColor: Colors.lightBlue,
          accentColor: Colors.pink,
          bottomAppBarColor: Colors.lightBlue,
          dividerColor: Colors.pink,
          focusColor: Colors.pink,
          highlightColor: Colors.pink,
          splashColor: Colors.pink,
          secondaryHeaderColor: Colors.pink,
          indicatorColor: Colors.pink,
          toggleableActiveColor: Colors.pink),
      routeInformationParser: _routeInformationParser,
      routerDelegate: routerDelegate,
      backButtonDispatcher: RootBackButtonDispatcher(),
    );
  }
}
