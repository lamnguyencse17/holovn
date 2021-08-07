import 'package:flutter/material.dart';
import 'package:holovn_mobile/models/schedule.dart';

class LiveCardDescription extends StatelessWidget {
  final Schedule schedule;
  LiveCardDescription(this.schedule);

  getAdaptiveTextSize(BoxConstraints constraints, dynamic value) {
    return (value / 600 ) * constraints.maxWidth;
  }

  @override
  Widget build(BuildContext context) {
    return LayoutBuilder(
        builder: (BuildContext context, BoxConstraints constraints) {
        return Center(
            child: FractionallySizedBox(
                widthFactor: 0.9, child: _wideDescription(constraints, schedule)));
    });
  }

  Widget _wideDescription(BoxConstraints constraints, Schedule schedule){
    return Row(
      children: [
        SizedBox(
          width: 0.20*constraints.maxWidth,
          height: 100,
          child: CircleAvatar(
              radius: 35,
              backgroundImage:
              NetworkImage(schedule.channel.photo, scale: 0.01,)),
        ),
        SizedBox(
            width: 0.7*constraints.maxWidth,
            height: 100,
            child: Column(
              children: [
                Align(
                  alignment: Alignment.centerLeft,
                  child: Text(schedule.title,
                      textAlign: TextAlign.start,
                      style: TextStyle(
                          fontSize:
                          getAdaptiveTextSize(constraints, 24))),
                ),
                Align(
                    alignment: Alignment.centerLeft,
                    child: Text(schedule.channel.name,
                        textAlign: TextAlign.start,
                        style: TextStyle(
                            fontSize:
                            getAdaptiveTextSize(constraints, 24)))),
                Align(
                    alignment: Alignment.centerLeft,
                    child: Text(schedule.startScheduled.toString(),
                        textAlign: TextAlign.start,
                        style: TextStyle(
                            fontSize:
                            getAdaptiveTextSize(constraints, 24)))),
              ],
            ))
      ],
    );
  }
}
