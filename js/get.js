var getTodos = function (callback) {
	"use strict;"

	// expects array of entries at URL path /todo
	$.getJSON("todo", function(entries) {
		if (callback && typeof callback === "function") {
			callback(entries);
		}
	});
};

var sortTodosBy = function (sortBy, entries) {
	"use strict;"

	if (!sortBy || typeof sortBy !== "string") {
		return [];
	}
	if (!entries || typeof entries !== "object" || !entries.hasOwnProperty("length")) {
		return [];
	}

	entries.sort(function(a, b) {
		if (sortBy === "Priority") {
			if (a.Priority !== "" && b.Priority !== "") {
				if (a.Priority < b.Priority) {
					return -1;
				} else if (a.Priority > b.Priority) {
					return 1;
				} else {
					return 0;
				}
			} else if (a.Priority !== "") {
				return -1;
			} else if (b.Priority !== "") {
				return 1;
			} else {
				return 0;
			}
		}
	});

	return entries;
};

var displayTodos = function (entries) {
	"use strict;"

	if (entries && entries !== null) {
		var todosDiv = $("#todos"),
			panel, panelHeading, panelFooter = null,
			heading = "", footer = "";

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

			if (element.hasOwnProperty("Description")) {
				$("<div>").attr("class", "panel-body").html(element.Description).appendTo(panel);
			}

			panelFooter.html(footer).appendTo(panel);

			panel.appendTo(todosDiv);
		})
	}
};