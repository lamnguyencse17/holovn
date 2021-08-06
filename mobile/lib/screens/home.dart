import 'package:flutter/material.dart';
import 'package:holovn_mobile/models/schedule.dart';
import 'package:holovn_mobile/models/schedule_list.dart';
import 'package:holovn_mobile/providers/schedule.dart';
import 'package:holovn_mobile/widget/home/live_card.dart';

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
  List<String> _stringList = [];
  late ScheduleList _scheduleList;
  @override
  void initState(){
    super.initState();
    WidgetsBinding.instance!.addObserver(this);
    fetchSchedule().then((fetchedSchedules) => {
      _scheduleList.values = fetchedSchedules.values
    });
    _stringList = ["a", "b", "c", "d"];
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
        body: GridView.count(
            crossAxisCount: 2,
            children: _stringList.map<Widget>(
                    (string) => Container(child: new LiveCard("title " + string, "https://yt3.ggpht.com/ytc/AKedOLSIXjxaKvTQAgBpIwZtdI_Ux_cUEi3wefTTSVZW1w=s800-c-k-c0x00ffffff-no-rj"))).toList()
        )
    );
  }
}