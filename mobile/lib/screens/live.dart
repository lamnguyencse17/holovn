import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:holovn_mobile/models/schedule.dart';
import 'package:holovn_mobile/router/router.dart';
import 'package:holovn_mobile/widget/live/player.dart';

class Live extends StatefulWidget {
  final Schedule? schedule;
  final String? liveId;
  Live(this.schedule, this.liveId);

  @override
  _LiveState createState() => _LiveState(schedule, liveId);
}

class _LiveState extends State<Live> {
  final Schedule? schedule;
  final String? liveId;
  final AppRouter router = Get.find(tag: "router");

  _LiveState(this.schedule, this.liveId);

  @override
  Widget build(BuildContext context) {
    var _schedule = this.schedule;
    if (_schedule == null) {
      var _liveId = this.liveId;
      if (_liveId == null) {
        return Container(child: Text("PENDING"));
      }
      router.navigate(null, null);
      return Container();
    } else {
      return Scaffold(
        appBar: AppBar(
          title: Text(_schedule.title),
        ),
        body: ListView(
          children: [Text(_schedule.title), Player(liveId!, _schedule.status)],
        ),
      );
    }
  }
}
