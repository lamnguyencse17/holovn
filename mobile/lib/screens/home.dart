import 'package:flutter/material.dart';
import 'package:holovn_mobile/models/channel.dart';
import 'package:holovn_mobile/models/schedule.dart';
import 'package:holovn_mobile/models/schedule_list.dart';
import 'package:holovn_mobile/providers/schedule.dart';
import 'package:holovn_mobile/screens/home/home_layout_builder.dart';
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
    // fetchSchedule().then((fetchedSchedules) => {
    //   _scheduleList.values = fetchedSchedules.values
    // });
    var channel = new Channel("UCp6993wxpyDPHUpavwDFqgg", "SoraCh. ときのそらチャンネル", "Hololive", "vtuber", "https://yt3.ggpht.com/ytc/AKedOLQO9Vyz7ysAwPSio5xvkw6n0xvlyDu7A9eawqIH3w=s800-c-k-c0x00ffffff-no-rj", "Tokino Sora");
    var schedule = new Schedule("QaKgoFyHPzg", "【ロマサガ3】カタリナ主人公で進めていきます！！【#ときのそら生放送】", "stream", DateTime.now(),  DateTime.now(),  DateTime.now(), 0, "upcoming", channel);
    var list = [schedule, schedule, schedule, schedule];
    _scheduleList = new ScheduleList(list);
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