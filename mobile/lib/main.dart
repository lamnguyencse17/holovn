import 'package:flutter/material.dart';
import 'package:holovn_mobile/router/route_parser.dart';
import 'package:holovn_mobile/router/router.dart';
import 'package:get/get.dart';
void main() {
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
    // routerDelegate.navigate(null, null);
  }
  RouteParser _routeInformationParser = RouteParser();

  @override
  Widget build(BuildContext context) {
    return MaterialApp.router(
        title: "Holovn - A Vietnamese Hololive Fan App",
        theme: ThemeData(primarySwatch: Colors.blue),
        routeInformationParser: _routeInformationParser,
        routerDelegate: routerDelegate);
  }
}
