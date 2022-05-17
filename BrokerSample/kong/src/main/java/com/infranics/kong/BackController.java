package com.infranics.kong;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;

@RestController
@RequestMapping("/search")
public class BackController {

	private static final Logger logger = LoggerFactory.getLogger(BackController.class);


	@RequestMapping(value="/get", method = {RequestMethod.POST, RequestMethod.GET} )
	public List<Data> search(Map model) {
		System.out.println("request /search/get");

		Data data = new Data();
		data.setId("1");
		data.setName("one");

		List<Data> list = new ArrayList<Data>();
		list.add(data);

		return list;
	}

	@RequestMapping(value="/getTwo", method = {RequestMethod.POST, RequestMethod.GET})
	public List<Data> searchTwo(Map model) {
		System.out.println("request /search/getTwo");

		Data data = new Data();
		data.setId("1");
		data.setName("one");

		List<Data> list = new ArrayList<Data>();
		list.add(data);

		data = new Data();
		data.setId("2");
		data.setName("two");

		list.add(data);

		return list;
	}

}

