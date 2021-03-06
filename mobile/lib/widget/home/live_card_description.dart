import 'package:flutter/material.dart';
import 'package:holovn_mobile/models/schedule.dart';
import 'package:holovn_mobile/router/router.dart';
import 'package:url_launcher/url_launcher.dart';
import 'package:get/get.dart';

class LiveCardDescription extends StatelessWidget {
  final Schedule schedule;
  LiveCardDescription(this.schedule);
  AppRouter router = Get.find(tag: "router");
  getAdaptiveTextSize(BoxConstraints constraints, dynamic value) {
    return (value / 650) * constraints.maxWidth;
  }

  String estimateSchedule(DateTime date) {
    var currentTime = DateTime.now();
    var hourDifference = date.difference(currentTime).inHours;
    var minuteDifference = date.difference(currentTime).inMinutes;
    if (hourDifference <= 0 && minuteDifference <= 0){
      return "Live at the moment";
    }
    if (hourDifference <= 0){
      return "Starts in " + minuteDifference.toString() + " minutes";
    }
    return "Starts in " + hourDifference.toString() + " hours";
  }

  launchChannelPage(String channelId) async{
    var url = "https://www.youtube.com/channel/" + channelId;
    if (await canLaunch(url)){
      await launch(url);
    } else {
      throw 'Could not launch $url';
    }
  }

  @override
  Widget build(BuildContext context) {
    return LayoutBuilder(
        builder: (BuildContext context, BoxConstraints constraints) {
      return Center(
          child: FractionallySizedBox(
              widthFactor: 0.9,
              child: _wideDescription(context, constraints, schedule)));
    });
  }

  Widget _wideDescription(BuildContext context, BoxConstraints constraints, Schedule schedule) {
    return Row(
      children: [
        SizedBox(
          width: 0.20 * constraints.maxWidth,
          height: 100,
          child: CircleAvatar(
              radius: 35,
              backgroundImage: NetworkImage(
                schedule.channel.photo,
                scale: 0.01,
              )),
        ),
        SizedBox(
            width: 0.7 * constraints.maxWidth,
            height: 170,
            child: Column(
              children: [
                Align(
                  alignment: Alignment.centerLeft,
                  child: Padding(
                      padding: EdgeInsets.fromLTRB(5, 5, 0, 5),
                      child: InkWell(onTap: () => router.navigate(schedule, schedule.id), child:Text(
                        schedule.title,
                        textAlign: TextAlign.start,
                        style: TextStyle(
                            fontSize: getAdaptiveTextSize(constraints, 24)),
                      ))),
                ),
                Align(
                    alignment: Alignment.centerLeft,
                    child: Padding(
                        padding: EdgeInsets.fromLTRB(5, 10, 0, 5),
                        child: InkWell(
                            onTap: () => launchChannelPage(schedule.channel.id),
                            child: Text(schedule.channel.name,
                                textAlign: TextAlign.start,
                                style: TextStyle(
                                    fontSize:
                                        getAdaptiveTextSize(constraints, 24) +
                                            3,
                                    fontWeight: FontWeight.w600))))),
                Align(
                    alignment: Alignment.centerLeft,
                    child: Padding(
                        padding: EdgeInsets.fromLTRB(5, 5, 0, 5),
                        child: Text(
                            estimateSchedule(this.schedule.startScheduled),
                            textAlign: TextAlign.start,
                            style: TextStyle(
                                fontSize: getAdaptiveTextSize(constraints, 24),
                                fontWeight: FontWeight.w600)))),
              ],
            ))
      ],
    );
  }
}
