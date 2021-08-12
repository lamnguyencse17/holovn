import 'package:flutter/material.dart';
import 'package:holovn_mobile/models/schedule.dart';
import 'package:holovn_mobile/providers/schedule.dart';
import 'package:holovn_mobile/screens/home/home_layout_builder.dart';
import 'package:holovn_mobile/widget/drawer_nav.dart';

class HomePage extends StatefulWidget {
  HomePage({Key? key, required this.title}) : super(key: key);

  final String title;

  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> with WidgetsBindingObserver {
  List<Schedule> _scheduleList = [];

  Future<void> refreshSchedules() async {
    var fetchedSchedules = await fetchSchedule();
    this.setState(() {
      _scheduleList = fetchedSchedules;
    });
  }

  _HomePageState();
  @override
  void initState() {
    super.initState();
    WidgetsBinding.instance!.addObserver(this);
    refreshSchedules();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        drawer: DrawerNav(),
        appBar: AppBar(
          title: Text(widget.title),
        ),
        body: HomeLayoutBuilder(_scheduleList, refreshSchedules));
  }
}
