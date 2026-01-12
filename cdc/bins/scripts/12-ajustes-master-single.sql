SET search_path TO bins;

-- adjust country codes for MasterCard matched bins
-- fonte: https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
update bins.bins bins
  set pais = case pais
               when '040' then 'AUT'
               when '076' then 'BRA'
               when '170' then 'COL'
               when '280' then 'DEU'
               when '380' then 'ITA'
               when '528' then 'NLD'
               when '840' then 'USA'
               else pais
end
where bandeira = 'MASTER2';