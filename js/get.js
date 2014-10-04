var getTodos = function () {
	"use strict;"

	// expects array of entries
	$.getJSON("todo", function(entries) {
		var todosDiv = $("#todos");
		var panel, panelHeading, panelFooter = null;
		var heading = "", footer = "";

		$.each(entries, function (index, element) {
			heading = "";
			footer = "";
			panelHeading = $("<div>").attr("class", "panel-heading");
			panelFooter = $("<div>").attr("class", "panel-footer");

			if (element.hasOwnProperty("Priority")) {

				if (element.Priority === "(A)") {
					panel = $("<div>").attr("class", "panel panel-danger");
				} else if (element.Priority === "(B)") {
					panel = $("<div>").attr("class", "panel panel-warning");
				} else if (element.Priority === "(C)") {
					panel = $("<div>").attr("class", "panel panel-info");
				} else if (element.Priority === "(D)") {
					panel = $("<div>").attr("class", "panel panel-success");
				} else {
					panel = $("<div>").attr("class", "panel panel-primary");
				}

				heading = element.Priority;
			}
			if (element.hasOwnProperty("Contexts")) {
				$.each(element.Contexts, function (index, context) {
					footer += " " + context;
				})
			}
			if (element.hasOwnProperty("Projects")) {
				$.each(element.Projects, function (index, project) {
					footer += " " + project;
				})
			}

			$("<h3>").attr("class", "panel-title").html(heading).appendTo(panelHeading);
			panelHeading.appendTo(panel);
			panelFooter.html(footer).appendTo(panel);

			if (element.hasOwnProperty("Description")) {
				$("<div>").attr("class", "panel-body").html(element.Description).appendTo(panel);
			}

			panel.appendTo(todosDiv);
		})
	});
};