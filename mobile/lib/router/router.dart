import 'package:flutter/material.dart';
import 'package:holovn_mobile/models/schedule.dart';
import 'package:holovn_mobile/router/route_path.dart';
import 'package:holovn_mobile/screens/home.dart';
import 'package:holovn_mobile/screens/live.dart';

class AppRouter extends RouterDelegate<RoutePath> with ChangeNotifier, PopNavigatorRouterDelegateMixin<RoutePath>{
  final GlobalKey<NavigatorState> navigatorKey;
  final _pages = <Page>[];

  Schedule? _selectedSchedule;
  String? liveId;
  bool show404 = false;

  @override
  Widget build(BuildContext context) {
    return Navigator(
      pages: [
        MaterialPage(
          key: ValueKey("Home"),
          child: HomePage(title: 'Holovn - A Vietnamese Hololive Fan App', navigateToLive: _navigateToSelectedSchedule),
        ),
        if (_selectedSchedule != null || liveId != null) MaterialPage(key: ValueKey("Live"),child: Live(_selectedSchedule, liveId))
      ],
      onPopPage: (route, result) {
        if (!route.didPop(result)) {
          return false;
        }
        _selectedSchedule = null;
        liveId = null;
        show404 = false;
        notifyListeners();
        return true;
      },
    );
  }
  //
  // @override
  // // TODO
  // GlobalKey<NavigatorState> get navigatorKey => throw UnimplementedError();

  @override
  Future<bool> popRoute() {
    // TODO: implement popRoute
    return super.popRoute();
  }

  @override
  Future<void> setNewRoutePath(RoutePath path) async {
    if (path.isUnknown == null) {
      _selectedSchedule = null;
      show404 = true;
      return;
    }
    if (path.isLivePage) {
      if (path.schedule == null) {
        show404 = true;
        return;
      }
    }
    _selectedSchedule = null;
    show404 = false;
  }

  RoutePath get currentConfiguration {
    if (show404) {
      return RoutePath.unknown();
    }

    return _selectedSchedule == null
        ? RoutePath.home()
        : RoutePath.live(_selectedSchedule, liveId);
  }

  void _navigateToSelectedSchedule(Schedule? schedule, String? scheduleId){
    _selectedSchedule = schedule;
    liveId = scheduleId;
    notifyListeners();
  }

  AppRouter() : navigatorKey = GlobalKey<NavigatorState>();
}