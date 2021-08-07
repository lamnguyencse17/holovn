import 'package:flutter/material.dart';
import 'package:holovn_mobile/models/schedule.dart';
import 'package:holovn_mobile/widget/home/live_card_description.dart';

class LiveCard extends StatelessWidget {
  final Schedule schedule;

  LiveCard(this.schedule);



  Widget build(BuildContext context) {
    return Center(
      child: FractionallySizedBox(
          widthFactor: 0.9,
          child: Card(
            child: Column(
              children: [
                new Expanded(
                    child: Image.network(
                        "https://i.ytimg.com/vi/" +
                            schedule.scheduleId +
                            "/sddefault_live.jpg",
                        fit: BoxFit.fitWidth)),
                LiveCardDescription(schedule)
              ],
            ),
          )),
    );
  }
}
