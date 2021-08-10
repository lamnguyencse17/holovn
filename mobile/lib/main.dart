import 'package:flutter/material.dart';
import 'package:holovn_mobile/router/route_parser.dart';
import 'package:holovn_mobile/router/router.dart';
import 'package:get/get.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:flutter/services.dart' show rootBundle;

Future main() async {
  try {
    // print(await rootBundle.loadString('assets/.env'));
    await dotenv.load(fileName: "assets/.env");
    // print(dotenv.env["HOLOVN_SERVER"]);
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
          // primaryColor: Colors.pinkAccent,
          accentColor: Colors.pinkAccent),
      routeInformationParser: _routeInformationParser,
      routerDelegate: routerDelegate,
      backButtonDispatcher: RootBackButtonDispatcher(),
    );
  }
}
