import 'package:flutter/material.dart';
import 'package:holovn_mobile/models/schedule.dart';

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
  _LiveState(this.schedule, this.liveId);
  @override
  Widget build(BuildContext context) {
    var _schedule = this.schedule;
    if (_schedule == null){
      return Container();
    } else {
      return Container(child: Text(_schedule.title),);
    }
  }
}