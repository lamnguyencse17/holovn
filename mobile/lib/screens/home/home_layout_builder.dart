import 'package:flutter/material.dart';
import 'package:holovn_mobile/models/schedule.dart';
import 'package:holovn_mobile/widget/home/live_card.dart';

class HomeLayoutBuilder extends StatelessWidget {
  final List<Schedule> scheduleList;
  HomeLayoutBuilder(this.scheduleList);

  @override
  Widget build(BuildContext context) {
    return LayoutBuilder(
        builder: (BuildContext context, BoxConstraints constraints) {
      if (constraints.maxWidth > 1000) {
        return _homeGridView(scheduleList, 3);
      }else if (constraints.maxWidth > 600) {
        return _homeGridView(scheduleList, 2);
      } else {
        return _homeGridView(scheduleList, 1);
      }
    });
  }

  Widget _homeGridView(List<Schedule> _scheduleList, int count) {
    return Center(
        child: FractionallySizedBox(child: GridView.count(
        crossAxisCount: count,
        children: _scheduleList
            .map<Widget>((schedule) => Container(child: new LiveCard(schedule)))
            .toList())));
  }
}
