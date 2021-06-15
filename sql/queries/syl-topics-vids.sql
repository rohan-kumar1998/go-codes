SELECT 
	syllabus.name, subject.name, Topic.name, subtopic.name, Concept.name, Video_segment.name, Video_segment.video_name 
	FROM 
		syllabus 
        INNER JOIN subject 
			ON syllabus.id=subject.syl_id AND syllabus.id=? AND syllabus.is_active=1 AND subject.is_active=1 
			INNER JOIN Topic 
				ON Topic.sub_id=subject.id AND Topic.is_active=1 AND subject.is_active=1
                INNER JOIN subtopic
					ON subtopic.top_id=Topic.id AND Topic.is_active=1 AND subtopic.is_active=1
					INNER JOIN Concept
						ON Concept.subtopic_id=subtopic.id AND Concept.is_active=1 AND subtopic.is_active=1
                        INNER JOIN Video_segment
						ON Video_segment.concept_id=Concept.id AND Video_segment.is_active=1 AND Concept.is_active=1; 