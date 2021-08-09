import 'package:flutter/material.dart';
import 'package:holovn_mobile/models/schedule.dart';
import 'package:holovn_mobile/providers/schedule.dart';
import 'package:holovn_mobile/screens/home/home_layout_builder.dart';

class HomePage extends StatefulWidget {

  HomePage({Key? key, required this.title}) : super(key: key);

  // This widget is the home page of your application. It is stateful, meaning
  // that it has a State object (defined below) that contains fields that affect
  // how it looks.

  // This class is the configuration for the state. It holds the values (in this
  // case the title) provided by the parent (in this case the App widget) and
  // used by the build method of the State. Fields in a Widget subclass are
  // always marked "final".

  final String title;

  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> with WidgetsBindingObserver {
  List<Schedule> _scheduleList = [];

  _HomePageState();
  @override
  void initState(){
    super.initState();
    WidgetsBinding.instance!.addObserver(this);
    fetchSchedule().then((fetchedSchedules) => {
      this.setState(() {
        _scheduleList = fetchedSchedules;
      })
    });
  }

  @override
  Widget build(BuildContext context) {
    // This method is rerun every time setState is called, for instance as done
    // by the _incrementCounter method above.
    //
    // The Flutter framework has been optimized to make rerunning build methods
    // fast, so that you can just rebuild anything that needs updating rather
    // than having to individually change instances of widgets.
    return Scaffold(
        appBar: AppBar(
          // Here we take the value from the MyHomePage object that was created by
          // the App.build method, and use it to set our appbar title.
          title: Text(widget.title),
        ),
        body: HomeLayoutBuilder(_scheduleList)
    );
  }
}