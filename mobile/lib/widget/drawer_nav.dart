import 'package:flutter/material.dart';

class DrawerNav extends StatelessWidget {
  GlobalKey<ScaffoldState> _scaffoldKey = new GlobalKey<ScaffoldState>();
  @override
  Widget build(BuildContext context) {
    return Drawer(
        child: ListView(padding: EdgeInsets.zero, children: [
      DrawerHeader(
        child: Text("HoloVN"),
        decoration: BoxDecoration(color: Colors.blue),
      ),
      ListTile(
        title: Text("Home"),
          onTap: () => _scaffoldKey.currentState!.openEndDrawer()
      ),
    ]));
  }
}
