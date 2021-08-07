import 'package:flutter/material.dart';
import 'package:holovn_mobile/router/route_parser.dart';
import 'package:holovn_mobile/router/router.dart';

void main() {
  runApp(HolovnApp());
}

class HolovnApp extends StatefulWidget {
  @override
  State<StatefulWidget> createState() => _HolovnAppState();
}

class _HolovnAppState extends State<HolovnApp> {
  void initState() {
    super.initState();
  }

  AppRouter _routerDelegate = AppRouter();
  RouteParser _routeInformationParser = RouteParser();
  @override
  Widget build(BuildContext context) {
    // return MaterialApp(
    //     title: 'Home Page',
    //     theme: ThemeData(
    //       primarySwatch: Colors.blue,
    //     ),
    //     );
    return MaterialApp.router(
        title: "Holovn - A Vietnamese Hololive Fan App",
        theme: ThemeData(primarySwatch: Colors.blue),
        routeInformationParser: _routeInformationParser,
        routerDelegate: _routerDelegate);
  }
}
