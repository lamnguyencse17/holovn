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
        return _homeGridView(scheduleList, 3, constraints);
      } else if (constraints.maxWidth > 600) {
        return _homeGridView(scheduleList, 2, constraints);
      } else {
        return _homeGridView(scheduleList, 1, constraints);
      }
    });
  }

  Widget _homeGridView(
      List<Schedule> _scheduleList, int count, BoxConstraints constraints) {
    return Center(
        child: SizedBox(
          width: constraints.maxWidth,
            height: constraints.maxHeight,
            child: DefaultTabController(
                length: 2,
                child: ListView(
                  children: [
                    TabBar(tabs: [
                      Tab(
                        text: "Live",
                      ),
                      Tab(
                        text: "Past",
                      )
                    ]),
                    SizedBox(
                      width: constraints.maxWidth * 0.9,
                      height: constraints.maxHeight,
                      child: TabBarView(children: [
                        GridView.count(
                            crossAxisCount: count,
                            children: _scheduleList
                                .map<Widget>((schedule) =>
                                    Container(child: new LiveCard(schedule)))
                                .toList()),
                        Container(
                          child: Text("PENDING"),
                        )
                      ]),
                    )
                  ],
                ))));
  }
}
